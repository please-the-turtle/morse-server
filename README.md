# morse-server
A service that returns Morse code in wav format.  
## ![GET](https://img.shields.io/badge/GET-green) /?m=hello_world&freq=1000&dotlen=200ms
Params:
- **m=_string_ — Requred, the maximum number of characters is 50.**
  String to convert to Morse code.
- **freq=_float_** — A float number in the range from 20.0 to 15000.0.
  Frequency of the dot sound in in hertz.
- **dotlen=_duration_** — A float number in the range from 20ms to 2000ms.
  The length of the sound of one point in milliseconds.
