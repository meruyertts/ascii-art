# ascii-art

Ascii-art is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII.  

A graphic representation using ASCII, is to write the string received using ASCII characters, as you can see in the example below:  

@@@@@@BB@@@@``^^``^^``@@BB$$@@BB$$
@@%%$$$$^^^^WW&&8888&&^^""BBBB@@@@
@@@@@@""WW8888&&WW888888WW``@@@@$$
BB$$``&&&&WWWW8888&&&&8888&&``@@@@
$$``&&WW88&&88&&&&8888&&88WW88``$$
@@""&&&&&&&&88888888&&&&&&88&&``$$
``````^^``^^^^^^````""^^``^^``^^``
""WW^^@@@@^^``````^^BB@@^^``^^&&``
^^&&^^@@````^^``&&``@@````^^^^&&``
``WW&&^^""``^^WW&&&&""``^^^^&&88``
^^8888&&&&&&WW88&&88WW&&&&88&&WW``
@@``&&88888888WW&&WW88&&88WW88^^$$
@@""88&&&&&&&&888888&&``^^&&88``$$
@@@@^^&&&&&&""``^^^^^^8888&&^^@@@@
@@@@@@^^888888&&88&&&&MM88^^BB$$$$
@@@@@@BB````&&&&&&&&88""``BB@@BB$$
$$@@$$$$$$$$``````````@@$$@@$$$$$$  
This project handles an input with numbers, letters, spaces, special characters and \n.  

Banner file(standard.txt) with a specific graphical template representation using ASCII is given. The file is formatted in a way that is not necessary to change them.  

# Banner Format

* Each character has a height of 8 lines.  
* Characters are separated by a new line \n.  
* Here is an example of ' ', '!' and '"'(one dot represents one space):  
......
......
......
......
......
......
......
......

._..
|.|.
|.|.
|.|.
|_|.
(_).
....
....

._._..
(.|.).
.V.V..
......
......
......
......
......

# Usage

student$ go run . "" | cat -e  
student$ go run . "\n" | cat -e  
$  
student$ go run . "Hello\n" | cat -e  
 _    _          _   _          $  
| |  | |        | | | |         $  
| |__| |   ___  | | | |   ___   $  
|  __  |  / _ \ | | | |  / _ \  $  
| |  | | |  __/ | | | | | (_) | $  
|_|  |_|  \___| |_| |_|  \___/  $  
                                $  
                                $  
$  
student$ go run . "hello" | cat -e  
 _              _   _          $  
| |            | | | |         $  
| |__     ___  | | | |   ___   $  
|  _ \   / _ \ | | | |  / _ \  $  
| | | | |  __/ | | | | | (_) | $  
|_| |_|  \___| |_| |_|  \___/  $  
                               $
                               $