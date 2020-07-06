# dplaylist

dplaylist is a free and open-source tool that allows you to generate a (YouTube) link from a list of video IDs/URLs contained in a file. That link will be used by YouTube to generate a playlist.

For various reasons, you may prefer not to use your Google account but you would still like to have a way to play and save your playlists. This tool makes it possible. The only thing you need to do is keep a file that contains one video link or ID per line.

## Use

1. Download the executable from the release section or build it yourself
1. Open a Terminal (cmd.exe on Windows)
1. Navigate to the location of the *dplaylist(.exe)* executable
1. Execute it by passing your text file as an argument: `./dplaylist hip-hop.txt`
1. The link is generated in the console

### Valid ID

Let's take an example: [Rick Astley - Never Gonna Give You Up (Video)](https://www.youtube.com/watch?v=dQw4w9WgXcQ). The ID of the video is effectively **dQw4w9WgXcQ**.

The following is an exhaustive list of the possible values that you may have in your text file:
- https://www.youtube.com/watch?v=dQw4w9WgXcQ
- https://youtu.be/dQw4w9WgXcQ
- dQw4w9WgXcQ

Do **NOT** put a blank line at the end. One line is one video.

## Build

1. Run `go build` or `go install`

### Cross-compile

Alternatively, you may wish to create an executable for Linux, MacOS (Darwin) and Windows.

If so, run `./cross-compile.sh`