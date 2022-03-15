import subprocess
import base64
import os
import re 
import json
from Crypto.Cipher import AES
from Crypto import Random
from Crypto.Util.Padding import pad, unpad
from flask import Flask, request, render_template, send_file, jsonify
from werkzeug.utils import secure_filename
import datetime

app = Flask(__name__)
app.config["UPLOAD_FOLDER"] = "./templates/uploads"

@app.route('/')
def index():
   return render_template('glass-portal.html')

@app.route('/create', methods = ['POST'])
def create():
    if request.method == 'POST':
        organisation = request.form.get('org_select1') #The organisation making this request

        f = request.files['file']

        if (os.path.isdir(app.config["UPLOAD_FOLDER"]) == False):
            os.mkdir(app.config["UPLOAD_FOLDER"])

        plaintext_file = os.path.join(app.config['UPLOAD_FOLDER'], secure_filename(f.filename))
        f.save(plaintext_file)

        starttime = datetime.datetime.now()
        encrypted_file, secret_key = encrypt_aes_cbc_256(plaintext_file)
        os.remove(plaintext_file) #delete original non-encrypted file

        #Generate CID and add encrypted file to IPFS.
        cid = add_to_ipfs(encrypted_file) 

        #record on hyperledger
        if (record_to_hyperledger(cid, secret_key, "/ipfs/" + cid, organisation)): #note that we only support creation of IPFS records for now hence the hardcoded /ipfs/ prefix.
            endtime = datetime.datetime.now()
            return jsonify({ "Status" : "Success", "Message" : "Entry encrypted and recorded on Hyperledger Fabric", "CID" : cid, "URI" : "/ipfs/" + cid, "Resource Name" : plaintext_file, "TimeTaken" : "{0}".format(datetime.datetime.strptime(str(endtime - starttime),'%H:%M:%S.%f').time())})

        endtime = datetime.datetime.now()
        return jsonify({ "Status" : "Fail", "Message" : "Failed to create resource. Either a duplicate exists or an exception occurred. Please try again.", "TimeTaken" : "{0}".format(datetime.datetime.strptime(str(endtime - starttime),'%H:%M:%S.%f').time())})
        
@app.route('/decrypt', methods = ['POST'])
def decrypt():
    if request.method == 'POST':
        f = request.files['file']
        ciphertext_file = os.path.join(app.config['UPLOAD_FOLDER'], secure_filename(f.filename))
        f.save(ciphertext_file)

        starttime = datetime.datetime.now() 
        secret = request.form.get('secret') #aes key/iv. Expected as hex based string value seperated by colon.
        decrypted_file = decrypt_aes_cbc_256(ciphertext_file, secret)
        endtime = datetime.datetime.now()

        #include the below line if we want to get the time taken for decrypt (for performance evaluation) instead of returning the actual decrypted file
        #return jsonify({"TimeTaken" : "{0}".format(datetime.datetime.strptime(str(endtime - starttime),'%H:%M:%S.%f').time())}) 

        return send_file(decrypted_file, as_attachment=True) #TODO: note that for now, we do not delete the decrypted file. In future work, it's better that the decryption takes place client side rather than server side. 

def add_to_ipfs(file):
    result = subprocess.run(["ipfs", "add", "{0}".format(file), "--quieter" ], stdout=subprocess.PIPE, text=True)
    return result.stdout.strip('\n')

def encrypt_aes_cbc_256(file):
    key = Random.get_random_bytes(32) #Generate AES key
    iv = Random.get_random_bytes(16) #Generate AES iv

    ct_bytes = None
    with open(file, "rb") as r:  
        data = r.read()        
        cipher = AES.new(key, AES.MODE_CBC, iv)
        ct_bytes = cipher.encrypt(pad(data, AES.block_size))

    encrypted_filename = file + "_encrypted"
    with open(encrypted_filename, "wb") as f:
        f.write(ct_bytes)

    return (encrypted_filename, "{0},{1}".format(key.hex(),iv.hex())) #Note key and iv are returned seperated by colon.

def decrypt_aes_cbc_256(file, secret):
    key = bytes.fromhex(secret.split(',')[0])
    iv = bytes.fromhex(secret.split(',')[1])

    pt_bytes = None
    with open(file, "rb") as r:  
        data = r.read()        
        cipher = AES.new(key, AES.MODE_CBC, iv)                
        pt_bytes = unpad(cipher.decrypt(data), AES.block_size)

        decrypted_filename = file + "_decrypted"
        with open(decrypted_filename, "wb") as f:
            f.write(pt_bytes)

    return (decrypted_filename)

def record_to_hyperledger(cid, key, uri, organisation):    
    #Encrypt content, generate CID, create entry in hyperledger fabric as below

    params = base64.b64encode(bytes('{{"CID":"{0}","key":"{1}","uri":"{2}"}}'.format(cid, key, uri), 'utf-8'))
    params = params.decode('ascii')
    result = subprocess.run(["./minifab", "invoke", "-p", '"createGlassResource"', "-t",  '{{"GlassResource" : "{0}" }}'.format(params), "-o", organisation ], stdout=subprocess.PIPE, text=True)

    #Naive approach to determine success. If any "error" text is observed, we simply assume the command failed for now.
    if ("error" in result.stdout):
        return False
    return True

@app.route('/read', methods=['POST'])
def read():
    organisation = request.form.get('org_select') #The organisation making this request. 
    cid = request.form.get('cid')

    starttime = datetime.datetime.now() 
    result = subprocess.run(["./minifab", "query", "-p", '"readGlassResourceKey","{0}"'.format(cid), "-t", "''", "-o", organisation], stdout=subprocess.PIPE, text=True)
    endtime = datetime.datetime.now() 

    result = extractResult(result.stdout)

    result["TimeTaken"] = "{0}".format(datetime.datetime.strptime(str(endtime - starttime),'%H:%M:%S.%f').time())
    return result

def extractResult(text):
    res = re.findall("{\"cid\":.*}", text)
    if (res):
        return json.loads(res[0])
    else:
        return json.loads( '{ "Status" : "Fail", "Message" : "The CID could not be found or read permission was denied." }' )
 

if __name__ == '__main__':
    app.run(debug = True)