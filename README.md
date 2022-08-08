```
  ______________                   _____            ________                  _____                       
  ___  __ \__  /______________________(_)___  __    __  ___/______________   ____(_)__________________    
  __  /_/ /_  __ \  __ \  _ \_  __ \_  /__  |/_/    _____ \_  _ \_  ___/_ | / /_  /_  ___/  _ \_  ___/    
  _  ____/_  / / / /_/ /  __/  / / /  / __>  <      ____/ //  __/  /   __ |/ /_  / / /__ /  __/(__  )     
  /_/     /_/ /_/\____/\___//_/ /_//_/  /_/|_|      /____/ \___//_/    _____/ /_/  \___/ \___//____/ 
                             {SSH Tool: Apart of the Pheonix suite}
 ─────────────────────────────────────────────────────────────────────────────────────────
```

# What is this tool

This tool is mainly built around people who one want a new way to execute ssh sessions, commands, and other things and for people who want to manage multiple sessions based on password authentication, Again this script was made for people who do NOT want to use private keys, why would you do that? no clue, but anyway yeah this script is pretty good for automating that.

This tool allows you to extract filepaths at an extreme amount of speed from a SSH server, send single files to a SSH server, execute commands, authenticate, ping, set new values, start shell sessions, and do much more all within one script session. This means you can sign onto one SSH server, logout, reset the values to what host and port you want to connect to and start a new shell session all without having to execute the program 

# Install

`git clone https://www.github.com/ArkAngeL43/Pheonix_SSH ; cd Pheonix_SSH ; go run main.go`

# Navigation with the shell

#### Console help menu ####

```
====== Module [ set ] 

Description: Set is a module that you will use to set session variables
 Instead of having to reload the script everytime you want to loginto another server
 You can rather just use the module set in the console to set the iformation, here is its syntax
 -------------------------------------------------------------------------------
 - set <pass/user/host/port/conn/auth/thei/tnam/twid/tmod> <value>
 : set pass <value>-> Will set the current password to use for executing and authenticating servers
 : set user <value>-> Will set the current username to use for executing and authenticating servers
 : set host <value>-> Will set the current hostname to use for executing and authenticating servers
 : set port <value>-> Will set the current port to use for connecting, pinging, and authenticating servers
 : set conn <value>-> Will set the connection method Example is `set tcp` which will connect using tcp
 : set auth <value>-> Will set the authentication type to use when starting a new ssh session
 : set thei <value>-> Will set the current shell hight when starting a new ssh session 
 : set tnam <value>-> Will set the current shell type when starting a new ssh session EG(linux, xterm)
 : set twid <value>-> Will set the current shell width when starting a new ssh session 
 : set tmod <value>-> Will set the current shell ECHO number mode when starting a new ssh session
 : set outf <value>-> Will set the output file for downloading files from a SSH server
 : set intd <value>-> Will set the input file / remote ssh destination filepath for file uploads
 
====== Module [ session ] 

Description: session is a module that you will use to auth, ping, or start a new
 SSH sesion or even log into remote SSH servers, this module requires no argument but rather
 preset values which can be initalized using the set module. 
 -------------------------------------------------------------------------------
 - session <auth/ping/start>
 : session auth  -> Will attempt to authenticate the current host with the passwords and usernames
 : session ping  -> Will attempt to ping a ssh server to see if it is alive, and output running a command
 : session start -> Will log into a SSH server and start a interactive command line based shell
 
====== Module [ console ] 

Description: Console based commands, usually used to clear output, load settings etc
 -------------------------------------------------------------------------------
 - console <cls/clear/settings>
 : console cls      -> Will clear all art, banners, and output from the screen
 : console clear    -> Same thing as console cls
 : console settings -> Will load all session variables (port, host, pass, file, key etc)
 
====== Module [ load ]

Description: load allows you to load customized configuration files and private
 SSH key files for further commands
 -------------------------------------------------------------------------------
 - load  <pkeyfile/infofile> <file location or path>
 : load infofile  -> Will load the Config.yaml path with custom settings and use that data for commands
 : load pkeyfile  -> Will load and set the private ssh key file for sessions
 
====== Module [ utils ] 

Description: sub utilities and tools that are designed to make the ssh experience better
 -------------------------------------------------------------------------------
 - utils <downloadf> <value>
 : utils dloadf <filepath>  -> Will download all files to a given / set output directory from a given server directory eg `ssh/pi/home/desktop` and output all files found

 : utils uloadf <filepath>  -> Will attempt to upload a single file to the remote SSH serverfrom a given source directory eg `home/desktop/file.txt`
 : utils uloadf <filepath>  -> NOTE: for `utils uloadf` the destination filepath `set intd` must also be a directory + filename /ssh/pi/whatever/output.txt`
```

# Bugs in this program 

there are not alot of bugs there are more of certian things that dont work all the time, so one thing that does not work all the time will be sending files, for some reason this has been hard to really get, this can bug out but sometimes it can work also note you can only upload files not directories.




### Examples ###

> Loading configuration files 

take the following file as a file to load your settings

**Config.yaml**

```yaml
SSH_data:
    SSH_Password: "password123"
    SSH_Username: "user"
    SSH_Hostname: "10.0.0.174"
    SSH_Portnumb: "22"
    SSH_Dial_method: "tcp"
    SSH_TTY_Auth_TYPE: "password"
    SSH_TTY_Height: "40"
    SSH_TTY_Width: "80"
    SSH_TTY_ECHO_Mode: "0"
    SSH_TTY_Name: "vt100"
    # note this section can be a bit tricky so read below for more details 
    #
    # Chances are if you configure this you will be using utils options
    # in this shell / program, make sure the paths are right, and are found
    # if they are not found the program will not apply it to the session
    # variable list.
    #
    # when sending the file or setting a directory for the file to be sent to the filepath 
    # must be ON THE DESTINATION SSH SERVER like so 
    #
    #
    # SSH_Dst_Directory_For_File_Upload: "/home/ssh_server/whatever/filepath/output.txt"
    #
    #
    # if you are DOWNLOADING files from a SSH server src will be the filepath of the files 
    # you want to download from the SSH server. the DST will be the output of the downloaded files 
    # when downloading files you can use `set outf` to set the output destination of the downloaded 
    # files from the SSH server that you downloaded the files from
    #
    #
    #
    # NOTE: DOWNLOADS ARE SPECIFIED FOR DIRECTORIES NOT INDIVIDUAL FILES THAT IS IF YOU ARE USING OPTIONS `utils dloadf path/to/ssh_directory`
    # UPLOADS ARE SINGLE FILES ONLY THAT IS IF YOU ARE USING `utils uloadf path/to/ssh_directory`
    # 
    # EXAMPLE DOWNLOAD TO DIRECTORY
    #
    # SSH_Dst_Directory_For_File_Download: "/home/my_pc/Desktop/Files"
    #
    #
    #  Summary 
    #
    # SSH_Dst_Directory_For_File_Upload is the location the file will be saved and sent to on the server 
    #
    # Example usage -> utils uloadf /home/xea43p3x/Desktop/filename.txt 
    # when you load this yaml file, the SSH DST directory for file upload will be used as the endpoint for the file to go to and be copied to
    #
    # SSH_Dst_Directory_For_File_Download is the location where files will be downloaded from the SSH serber onto on your home computer
    SSH_Dst_Directory_For_File_Upload: "/path/to/file"
    SSH_Dst_Directory_For_File_Download: "/path/to/file"
#
#
# Side NOTE
#
# Configuration files can be in ANY path, they just MUST have the same structure and names
```

in the shell you can load this file by running 

`load infofile Config.yaml`

which will produce the output 

```
|+|> Config :: New SSH Password    ->  password
|+|> Config :: New SSH Username    ->  user
|+|> Config :: New SSH Hostname    ->  10.0.0.174
|+|> Config :: New SSH Port        ->  22
|+|> Config :: New SSH Ping Method ->  tcp
|+|> Config :: New SSH Auth Method ->  password
|+|> Config :: New SSH TTY Height  ->  40
|+|> Config :: New SSH TTY Height  ->  80
|+|> Config :: New SSH TTY ECHO Mode ->  vt100
|+|> Config :: New SSH TTY Name    ->  vt100
|+|> Config :: New SSH DST File for upload    ->  /some/filepath
```

> setting a user

`set user username124`

> setting a password 

`set pass 122333434534543`

> setting a hostname

`set host 127.0.0.1`

> setting a port

`set port 22`

> setting a connection method as tcp

`set conn tcp`

> setting a directory to download files to on your computer 

`set outf desktop/home/computer`

> Setting a local file to upload 

`set intd temote/destination/filename.txt`

> Clearing the screen 

 `console cls` or `console clear`

> to view set variables

`console settings`

> to download 

`utils dloadf remote/ssh/directory/to/Download`

> to upload 

`utils uloadf file/to/upload/on/your/computer/filename.txt`

> Start a interactive shell 

`session start`

> Ping a ssh server 

`session ping`

> Auth a ssh server 

`session auth`

