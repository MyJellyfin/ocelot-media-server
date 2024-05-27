# Ocelot Media Server [W.I.P.]

### Purpose
A free and open-source alternative to Plex Media Server that should look *beautiful*

### Goal
Reach full feature parity to Plex and possibly even Jellyfin.

- [ ] Users
    - [X] Login with email and password
    - [X] Create an admin user 
    - [ ] Create secondary users
        - [ ] Have customized permissions for each user
- [ ] Encoding (ffmpeg)
    - [X] Generate .m3u8 file
    - [ ] Generate .ts files dynamically 
        - [ ] Thumbnails
        - [ ] Subtitles
- [ ] Library
    - [X] Add directories
    - [ ] Scan and add metadata for directories
        - [X] Scan directory for files
        - [ ] Match scanned files for metadata
            - [ ] Search metadata from external API's
        - [ ] Scan all folders on startup for changes
        - [ ] Detect changes while server is still running
    - [ ] Control permission for which users have access to which libraries
- [ ] Allow custom plugins
    - [ ] Themes
- [ ] Logging
    - [ ] Show user data (such as who's watching etc.)
    - [ ] Analytics

### Questions
Here I'll answer some questions about the project preemptively

#### What is the project status?
As of right now, the project is in an **extremely** early status. The need for both a client and a server to be built together means that this will be an extremely long process. Not to mention that I am a college student and don't have too much free time to work on this. 

#### What devices will the client support?
The client is being worked on using the cross-platform framework, [Flutter](https://flutter.dev/), which will allow the client to be easily deployed on a variety of devices. The platforms I am aiming to target and support officially are: **iOS, Android, MacOS, Windows, Linux, LG WebOS, AndroidTV, and Apple's tvOS.** This is not a finalized list and I may choose to drop support if it proves to be more troublesome than it's worth.  

#### Why not fork Jellyfin, a project that intends to have the same goal?
During my personal experience moving from Plex to Jellyfin I encountered a few hurdles and friction that left me less than satisfied.
Not to say Jellyfin isn't an amazing project because it is and will be probably superior to this project for a very long time, but I would like to create something aesthetically pleasing, fast, and written in Golang.

#### Then why not just make a client for Jellyfin? 
I have worked on and tried to build a client for Jellyfin, which has taught me a lot about how media streaming but ultimately I believe I can make a viable alternative free of the tech debt from Emby/Jellyfin. Also, I don't have a lot of experience with C# and would personally like an alternative built in Go. 


##### Development started on April 29, 2024, by Siddhant Madhur

