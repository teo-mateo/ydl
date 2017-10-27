# ydl
ydl - yt downloader

For my own personal use. 
I use Youtube a lot to listen to music and often I want to dl the songs and listen to them offline in my car, from an USB stick. 
I dislike youtube downloaders because they are breaking my flow.

So I made this downloader that is running on one of my VMs. 
It has a chrome extension that lets me just click a button when I like a song, and it will be downloaded and converted to MP3 on my server; then I only need to fetch them from time to time. 


# instructions:
1. git clone git@github.com:teo-mateo/ydl.git

*you only need the folder chrome-ext.

2. make a small change to a js file. This will implicitly create your profile.
   The file is **chrome-ext/popup.js**
   
   Modify this code:
   
    //AJAX call   
    var x = new XMLHttpRequest();   
    x.open('GET', "http://localhost:8080/ydl?who=otheruser&v="+url);   
    
    ----> change the **"who"** parameter from "otheruser" to however you want to name your user.
    ----> change the url from localhost to **http://bardici.ro:8080/...**
   
   save the file.
     -Cpt. Obvious
   
 3. install the extension
    Use Chrome
    
    Navigate to chrome://extensions/
    
    Click on "Load unpacked extension"
    
    Select the /chrome-ext folder
    
 
 4. my pretty face will show up in your chrome toolbar. click it when you like dat song.
 
 6. don't abuse it; don't delete other people's songs. 
    the system is being monitored 24/7 by myself, but only when I remember to do that. 
    suspicious activity will be thoroughly ignored.   
 
 7. don't be mean to puppies.
