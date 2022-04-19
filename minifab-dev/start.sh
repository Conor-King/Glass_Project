#!/bin/bash

mainselection=99
useless=0
tag=""
inputA=""
inputB=""
command1=""
command2=""
entityA=""
entityB=""
step1=false

while : 
do
  case "$mainselection" in
  #every number coresponds to option
  #exit
  0)
    echo "Exiting..."
    exit
    ;;

  
    1)
    cd ipfs
    go run main.go 2 >init.txt
    tag=$(tail -n 1 init.txt | tr -d '[]')
    myarray=($tag)
    # inputA is what is goin to be inputted in A entity
    # change the loop values if you want to include more elements in A or B Default: first 2 = A; last 2 = B
  printf "Specify a name for the first entity of the hyperledger network (ex. \"a\"): "
  read entityA  
  printf "\nSpecify a name for the second entity of the hyperledger network (ex. \"b\"): "
  read entityB  
  echo "initialising IPFS...."

    
    
    for a in {0..1..1}; do
      inputA="$inputA ${myarray[$a]}"
    done

  printf "\nState: \n -Entity: \"$entityA\" \n -Assets:$inputA \n"
    for b in {2..3..1}; do
      inputB="$inputB ${myarray[$b]}"
    done

  printf "\nState: \n -Entity: \"$entityB\" \n -Assets:$inputB \n \n"
  
  command1=" '\"init\",\"$entityA\",\"${inputA:1}\",\"$entityB\",\"${inputB:1}\"'"    
  cd ..
  command2="minifab initialize -p$command1"
  #echo "($command2)"
  #echo "$(pwd)"    
  printf "Make sure that you approved and commited the new chaincode before continuing to step 2!\n"
  echo "Press enter to continue..."
  step1=true
  #eval $command2
  read useless  
    mainselection=99
    ;;

  2)
    

    if [ $step1 = true ] 
    then
    printf "\nTo be imported in the HyperLedger: \n -Entity: \"$entityA\" \n -Assets:$inputA \n"
    printf "\nTo be imported in the HyperLedger: \n -Entity: \"$entityB\" \n -Assets:$inputB \n \n"
    printf "*** Make sure that: *** \n-minifab is up and running \n-the new chain-code is just installed \n-the new chaincode is commited and approved"


    printf "\n\nProceed? \n \"y\" to confirm \"n\" to cancel: "
    read useless
    if [ $useless = "y" ]
    then
   
    eval $command2
    else 
    mainselection=99
    fi
    
    else
    printf "\n \n Complete step 1 first! \n Press Enter to continue... "
    read useless  
  fi
    mainselection=99
    ;;

   3)
  # 	echo "Installing chaincode"

      printf "\nThe current format of the IPFS files is: \n"
      printf "Name, Address, Nino Number \n"

      cd ipfs
      go run main.go 1 

       echo "Press enter to continue..."
       read useless
       mainselection=99
   		;;

  
    99)
    printf "\n Make sure IPFS Daemon is running! \n\n"

    echo "==================================="
    echo "        ---- Options  ----"
    echo "==================================="
    echo "    \"1\" Initialise IPFS"
    echo "    \"2\" Write to the Hyperledger"
    echo "    \"3\" Create an Asset and Upload it to IPFS "
    echo "    \"0\" Exit"
        printf "\nInput: "

    read mainselection
    ;;

  esac
done
