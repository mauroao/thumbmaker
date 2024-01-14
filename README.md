# Thumbsmaker
 
## Links

https://devicetests.com/extract-frames-fps-scaling-ffmpeg

## Ideas

Suggestion name for program: "Thumbs Grid Video Maker" 

## Requirements

- Each cell in the sheet MUST be at least 300px wide;
- We suggest 3x9 (27 pics) or 2x10 (20 pics) and 900-1200px wide; 

## More stuff 

- qtd pictures (pics): 27 
- video duration (duration): 948 

> Time interval = duration / pics

Time interval = 948 / 27 = 35.11

## Tests

```
ffmpeg -i in.mp4 -vf "fps=(1/35.11),scale=400:-1" frames/%04d.jpeg
```
- "1/35.11" means that I want one screenshot at every 35.11 seconds. 
- "400:-1" means that I want 400 width, and height must be calculatet in the way that output maintain original scale.

## Get video duration

```
ffprobe -v error -show_entries format=duration -of default=noprint_wrappers=1:nokey=1 your_video_file.mp4
```
