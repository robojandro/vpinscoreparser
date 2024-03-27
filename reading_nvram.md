# Williams

Williams system pinball machines:
https://pinwiki.com/wiki/index.php/Williams_System_3_-_7#System_7

```
System 6
...
Gorgar  12-1979
Firepower   02-1980
Scorpion    07-1980
...

System 7
Black Knight    11-1980
Jungle Lord     02-1981
Pharaoh     05-1981
Solar Fire  07-1981
Barracora   09-1981
Warlok  10-1982
Defender    12-1982
Cosmic Gunfight     06-1982
```

## System 7

Max score 9,999,999

Black Knight and Barracora seem to start player 1's high score at the 102nd byte:

102 byte [11110000] = str [ð] = hex [f0] = number [240] = type uint8
103 byte [11110000] = str [ð] = hex [f0] = number [240] = type uint8
104 byte [11110001] = str [ñ] = hex [f1] = number [241] = type uint8
105 byte [11110001] = str [ñ] = hex [f1] = number [241] = type uint8
106 byte [11110001] = str [ñ] = hex [f1] = number [241] = type uint8
107 byte [11110001] = str [ñ] = hex [f1] = number [241] = type uint8
108 byte [11110001] = str [ñ] = hex [f1] = number [241] = type uint8
109 byte [11110001] = str [ñ] = hex [f1] = number [241] = type uint8
110 byte [11110001] = str [ñ] = hex [f1] = number [241] = type uint8
111 byte [11110000] = str [ð] = hex [f0] = number [240] = type uint8
112 byte [11110111] = str [÷] = hex [f7] = number [247] = type uint8

Bytes 111 and 112 together are the count of coins inserted for a maximum of 99 credits.

## System 6

Scorpion (1980), rom:scrpn_l1.nv, display tops out 999,999

In this case, the  hiscore was manually set to 555,555 and there was 1 credit in the system.  It's possible that the 6 refers to the number of digits used in score and there appears to be 4 padding characters between this score and the 01 credits.

According to:
    https://pinside.com/pinball/machine/firepower/scores
    https://pinside.com/pinball/machine/scorpion/scores

can do millions - so 9.999.999
    even though display for scorpion only has 

71  hex [f6] = number [246]
72  hex [f5] = number [245] .
73  hex [f5] = number [245] .
74  hex [f5] = number [245] .
75  hex [f5] = number [245] .
76  hex [f5] = number [245] .
77  hex [f5] = number [245] .
78  hex [f0] = number [240]
79  hex [f0] = number [240]
80  hex [f0] = number [240]
81  hex [f0] = number [240]
82  hex [f0] = number [240]
83  hex [f1] = number [241]

Firepower frpwr_b7.nv 1980-02 

Score is 87850 with 06 credits inserted

69  hex [f0] = number [240]
70  hex [f1] = number [241]
71  hex [f8] = number [248]
72  hex [f0] = number [240] 
73  hex [f0] = number [240] 1 
74  hex [f8] = number [248] 2
75  hex [f7] = number [247] 3 
76  hex [f8] = number [248] 4 
77  hex [f5] = number [245] 5 
78  hex [f0] = number [240] 6
79  hex [f0] = number [240] x
80  hex [f0] = number [240] x
81  hex [f0] = number [240] x
82  hex [f0] = number [240] 1
83  hex [f6] = number [246] 2

Actually the byte at index 71 might not be a limit to the score as this isn't backed up by the Firepower parsed output.

# Stern Electronics

The original Stern company, Stern Electronics made:

Seawitch (1980)  Stern M-200 MPU
Quicksilver (1980)
Flight 2000 (1980) 
Lightning (1981)  Stern M-200 MPU
Viper (1981)  Stern M-200 MPU
Iron Maiden (1982)  Stern M-200 MPU
  all used the Stern M-200 MPU Main Processing Unit
    which seem to use the U8 5101 RAM same as Bally

## Seawitch

Current high score is 180340 and there are 2 credits in the machine.

Max score is appearently 9,999,999

203  hex [18] = number [24]
204  hex [3] = number [3]
205  hex [40] = number [64]
206  hex [2] = number [2]

*score*
byte 202 is 1,000,000s
byte 203 is 10,000s
byte 204 is 100s
byte 205 is 1s
    a 10x multiplier each time

*credit*
byte 206 is 1s

## Viper 

Score is 98,010 and 3 credits in the system

203  hex [9] = number [9] = string [    ]
204  hex [80] = number [128] = string [<80>]
205  hex [10] = number [16] = string [^P]
206  hex [3] = number [3] = string [^C]

All these bytes are consistent between the 2 games.

And oddly when there are no credits the output following the score was this:

206  hex [0] = number [0] = string []    
207  hex [0] = number [0] = string []    
208  hex [0] = number [0] = string []    
209  hex [7] = number [7] = string []    
210  hex [0] = number [0] = string []    
211  hex [0] = number [0] = string []    
212  hex [1] = number [1] = string []    
213  hex [0] = number [0] = string []    
214  hex [0] = number [0] = string []    
215  hex [7] = number [7] = string []    


# Bally

https://www.ipdb.org/search.pl?searchtype=advanced&mpu=18

In the 1980s Bally made:

Silverball Mania (1978)
Flash Gordon (1980)
Skateball (1980)
Spectrum (1980)
Xenon (1980)
Fathom (1981)
Elektra (1981)
Vector (1981)
Medusa (1981)
Genesis (1982)
Dungeons & Dragons (1987)

## Elektra

System: Bally MPU AS-2518-35

Current high score 885050

looking for 
8
85

Bally is read in reverse!!!

With the number as the first part of the 

In a hex reader we see the score stored at:

                             offset      171
000160   FF FF FF FF FF FF 0F 5F 0F 5F 8F 8F FF 4F 0F FF FF FF FF

But that is not the order that the bytes print out in, length of output is 302 bytes.

302-171=131, look at position 131

The pattern 8F is not seen in the scanned bytes so perhaps I need to read it in backworks via the bufio scanner in Golang.

"What you need is a reverse buffered reader that initializes from a ReaderSeeker."

166  hex [f] = number [15] = string []
167  hex [5f] = number [95] = string [_]
168  hex [f] = number [15] = string []
169  hex [5f] = number [95] = string [_]
170  hex [8f] = number [143] = string []
171  hex [8f] = number [143] = string []

But when read in reverse it is position 131 that it starts at.

Current high score 885050

When we flipped the reading order we get:

171  hex [8f] = number [143] = string [] 
170  hex [8f] = number [143] = string [] 
169  hex [5f] = number [95] = string [_] 
168  hex [f] = number [15] = string []   
167  hex [5f] = number [95] = string [_] 
166  hex [f] = number [15] = string []   


## Fathom

Current score is 135,650 with at least 3 credits

In the hex editor, it also starts at position 171.

171  hex [1f] = number [31] = string []
170  hex [3f] = number [63] = string [?]
169  hex [5f] = number [95] = string [_]
168  hex [6f] = number [111] = string [o]
167  hex [5f] = number [95] = string [_]
166  hex [f] = number [15] = string []

7 digits max - 7,777,777

so really out read starting positions is at byte 170 and read until 166.




In the 1990's Bally made
 - Attack from Mars
 - Addams Family
 - Twilight Zone
 - Circus Voltaire
 - Theatre of Magic



----

### Stern reference:

https://en.wikipedia.org/wiki/Stern_(game_company)

https://pinwiki.com/wiki/index.php/Bally/Stern#Continuity_Chart_for_U8_5101_RAM_in_a_Bally_AS2518-17.2C_-35_or_Stern_MPU-100_Board

```
LEARN MORE: How does a 128 byte 5101 RAM occupy 256 memory locations?

If you look at a pinout of the 5101 memory, you will notice it is a 128 byte device. Yet, it is addressed by the MPU via 256 memory locations ($200-$2FF). This is because the 5101 is actually a 256 nibble device - a nibble is a half-byte (4 bits). So data stored to a 5101 in a pinball machine actually only stores half of the data byte being sent to it. Which half it stores is dependent on the board design. Bally and Stern use the upper nibble for storage, and Williams used the lower nibble. Stern MPU-200 boards have an additional 5101 at U13. This stores the lower nibble in conjunction with U8 storing the upper nibble of a byte saved to $200-$2FF. This allows MPU-200 games to store more data, and avoid doing some fancy processing by getting the data in and out of the non-volatile ram area.

For example, here's some pseudo-machine code for what happens:

    LOAD #$24 (the data you want to store is 24)
    STORE $231 (you want to store the data 24 at memory location $231)
    READ $231 (you want to read back the data you just stored)

The data returned is not #$24 as expected, but rather #$2F. The lower nibble was never stored, because the 5101 memory does not store data as bytes but rather as nibbles. To store #$24 properly would require splitting the byte into its nibbles '2' and '4'. The 2 would be stored in one memory location, while the 4 would be shifted, and stored in another memory location.

Showing the byte as binary might be helpful to visualize what's involved. The hex #$24 in binary is %00100100. Split into nibbles is %0010 (the 2) and %0100 (the 4). The upper nibble is the one the 5101 is able to store directly, but the position in the byte of the lower nibble prevents it from being stored. A shift operation is performed 4 times on the byte to reposition the lower nibble as the upper nibble, which enables it to be stored to the 5101. Each shift moves the binary pattern to the left one bit - here's the full sequence:

    Start=%00100100
    Shift left=%0100100x
    Shift left=%100100xx
    Shift left=%00100xxx
    Shift left=%0100xxxx

This gives the #$4 in the high nibble. All byte data has to be split this way to be saved, and recombined upon reading from boards (Bally -17 and -35, Stern MPU-100) with a single 5101 RAM chip. You can see why Stern added the second 5101 RAM to their boards. It makes programming much easier!
```

