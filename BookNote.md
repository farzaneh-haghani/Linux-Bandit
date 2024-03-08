# Chapter 1: COMPUTER CONCEPTS

#### Count of combination for unique binary numbers:

- Count of combination for n bits is `2 ^ n`
  - b = bit & B = byte
  - Combination for 4b or nibble: max is 1111 so the count will be 2 ^ 4 = 16
  - Combination for 1B or 8b: max is 1111 1111 so the count will be 2 ^ 8 = 256

### SI Prefixes (metric) || IEEE Prefixes (binary)

- `1K (Kilo) = 10 ^ 3` || `1Ki (Kibi) = 2 ^ 10 = 1024`
- `1M (Mega) = 10 ^ 6` || `1Mi (Mebi) = 2 ^ 20`
- `1G (Giga) = 10 ^ 9` || `1Gi (Gibi) = 2 ^ 30`
- `1T (Tera) = 10 ^ 12` || `1Ti (Tebi) = 2 ^ 40`

#### Binary prefixes vs metric prefixes:

- When dealing with `storage capacity`, such as the size of a file or the capacity of a storage device, it's common to use `binary prefixes` , where `megabyte = 2 ^ 20 bytes`.
  This is because computers use binary (base-2) numbering systems, and storage capacities are typically measured in powers of 2.

- When dealing with `Data Transfer Rates`, internet service providers and networking equipment often use `metric prefixes`, where `megabyte per second = 10 ^ 6 bytes per second`. This is based on the standard metric system (base-10) and is commonly used to express data transfer rates over networks.

### Hexadecimal

We use hexadecimal just to help us in convert decimal to binary faster because it is short.

2,320,695,04 in decimal = 1000 1010 0101 0010 1111 1111 0000 0000 in binary = 8A52FF00 in hexadecimal

<br/>
<br/>

# Chapter 2: BINARY IN ACTION

- `encoding`: Translating data into a digital format (0,1)
- `decoding`: Interpret the digital data

### Standards

- `ASCII` (American Standard Code for Information Interchange) handles 128 characters for English stored as a full byte (Although 7 bits was enough, it is with an extra bit, left as 0).
- `Unicode` handles characters used in nearly all languages.

### Colors and Images

- `Color` in digital using `RGB` approach which is 24 bits to represent Red, Green, and Blue to make a single color.

  - 8 bits for each color from 00000000 (0 decimal or 0 hex) to 11111111 (255 decimal or FF hex)

- `Image` in digital constructed by a grid of `pixels` on a two-dimensional plane.

  - Pixels is squares with a particular color.

  - `Bitmap images` store the RGB color data for each individual pixel.

  - `JPEG` and `PNG` use compression techniques to reduce the number of bytes required to store an image, as compared to a bitmap.

### Interpreting Binary Data

- A binary value like 011000010110001001100011 can be:
  - ASCII text string “abc.”
  - a 24-bit RGB color
  - a positive integer 6,382,179 in decimal.
- It all depends on the program is written to expect data in a particular format:
  - `Text editor program` may assume the data is text.
  - `Image viewer` may assume it is the color of a pixel in an image.
  - `calculator` may assume it is a number.
