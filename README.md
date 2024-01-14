# thumbmaker

## Prompt

Create a program using Go.
This program will use ffmpeg to create a thumbnail image grid of a video file.
The program must have a function create_thumbs that receive theese parameters:
- input_file: path to the video file, default is input.mp4;
- output_file: path to image to be created, default is out.png, or out.jpg (you are free to choose the best image extension);
- columns: a number of columns, default is 4;
- thumb_width: the width in pixels of every thumb image in the grid, default is 300; 
- interval: a number of seconds between every video screen capture, default is 30 seconds;
The main function must call create_thumbs function with default parameters.
The thumb height must be calculated based on the original video resolution.
 
## Links

https://devicetests.com/extract-frames-fps-scaling-ffmpeg

## Tests

ffmpeg -i in.mp4 -vf "fps=0.0167,scale=400:-1" frames/%04d.jpeg

## Other stuff

Thumbs Grid Video Maker
