<?php
$host = "127.0.0.1";
$port = "9872";
 
$socket = socket_create(AF_INET, SOCK_STREAM, SOL_TCP)
  or die("Unable to create socket\n");
 
@socket_connect($socket, $host, $port) or die("Connect error.\n");
 
if ($err = socket_last_error($socket))
{
 
  socket_close($socket);
  die(socket_strerror($err) . "\n");
}
include ('../message/pb_message.php');
include ('./pb_proto_test_new.php');



// generate message with the new definition with surname
// now just test the classes
$book = new AddressBook();
$person = $book->add_person();
$person->set_name('Nikolai');
$person = $book->add_person();
$person->set_name('Kordulla');
$person->set_surname('MySurname');

$phone_number = $person->add_phone();
$phone_number->set_number('0711');
$phone_number->set_type(Person_PhoneType::WORK);

$phone_number = $person->add_phone();
$phone_number->set_number('0171');
$phone_number->set_type(Person_PhoneType::MOBILE);

$phone_number = $person->add_phone();
$phone_number->set_number('030');

// serialize
$string = $book->SerializeToString();
for ($i=0;;$i++){
	$leng=strlen($string);
	$leng=pack("N",$leng);
	socket_write ($socket ,$leng,strlen($leng));
	$first=substr($string,0,2);
	$second=substr($string, 2);
	$l1=socket_write ($socket ,$first,strlen($first));
	sleep(rand(0,5));
	$l2=socket_write ($socket ,$second,strlen($second));
	echo $l2+$l1;
	sleep(rand(0,5));
}
socket_close($socket)
?>