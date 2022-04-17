#!/bin/sh



mainselection=99
useless=0

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
    echo "initialising IPFS...."
    cd ipfs
    go run main.go init
    echo "Press enter to continue..."
    read useless
    mainselection=99
		;;

	2) 
	echo "initialise ipfs"
    go run main.go input
    echo "Press enter to continue..."
    read useless
    mainselection=99
		;;
3) 
	echo "Installing chaincode"


    echo "Press enter to continue..."
    read useless
    mainselection=99
		;;
    99)
echo "\n Make sure IPFS Daemon is running! \n"

echo "==================================="
echo "        ---- Options  ----"
echo "==================================="
echo "Input:"
echo "      \"1\" Initialise IPFS"
echo "      \"2\" Input manually a file in IPFS"
echo "      \"3\" Install custom chaincode "
echo "      \"0\" to exit the menu"
read mainselection

esac
done 
