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
