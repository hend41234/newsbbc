# NewsBBC Project


This application uses a 3rd party, namely Apify.
Please create an account if you don't already have one.

### Setup

1. Create a `.env` file in the root directory of your project with the following content:

    ```.env
        BASE_URL_APIFY= "https://api.apify.com/v2"
        URL_NEW_DATASET_APIFY="https://api.apify.com/v2/acts/shu8hvrXbJbY3Eb9W/runs"
        TOKEN_APIFY="your_token_apify"
        #GRAPH API
        BASE_URL_GRAPH_API="https://graph.facebook.com/v22.0"
        USER_ID="your_user_id"
        SECRET_APP="your_secret_app"
        APP_ID="your_app_id"
        SYSTEM_USER_ACCESS="your_system_user_token"
    ```

2. Replace `all requirements` with your actual API key obtained from Meta and Apify.

### Getting a Meta API Key

* To get an API key from Meta, follow these steps:

    1. Go to the [Meta for Developers](https://developers.facebook.com/) website.
    2. Log in with your Facebook account.
    3. Navigate to the "My Apps" section and click on "Create App".
    4. Follow the prompts to create a new app.
    5. Once your app is created, go to the "Settings" > "Basic" section to find your App ID and App Secret.
    6. In the "Products" section, add the necessary products (e.g., Facebook Login, Graph API) to your app.
    7. Generate an access token in the "Tools" > "Access Token Tool" section.
    8. Use this access token as your API key in the `.env` file.
    9. Then go to business settings, and customize user pages and systems.

* For more detailed instructions, refer to the [Meta Developer Documentation](https://developers.facebook.com/docs/).



# Instructions for Use :

* ### `uploadreels.RunUploadsReels()` can be used to upload videos to all your pages, each video is only valid for 1 page,
    if you want 1 video for all pages you can change it in the section:

    * `app/upload_reels/run_uploads.go`
        ```go
                // Upload Video
        		// open detail_video.json
        		videoList := download.OpenDetailVideo() // moved code
        		detailVideo := videoList[0]             //moved code
        		runUpload := uploadphase.UploadPhaseModels{
        			Rupload:     getSession,
        			AccessToken: list.AccessToken,
        		}
        		filePath := detailVideo.NameVideo              // moved code
        		uploader, _ := runUpload.UploadPhase(filePath)
        		if !uploader.Success {
        			log.Fatal("upload error")
        		}
        		// published
        		publishParams := publish.PublishModelsParams{
        			EndPoint:    fmt.Sprintf("%v/%v/video_reels", app.Utils.BaseUrlGraphApi, list.ID),
        			AccessToken: list.AccessToken,
        			VideoID:     getSession.VideoID,
        			UploadPhase: "finish",
        			VideoState:  "PUBLISHED",
        			Description: detailVideo.Description,
        		}
        		publishParams.Publish()                                 
        		rm := download.RemoveIndexZeroDetailVideo(videoList)     // moved code
        		download.WriteJSON(rm)                                   // moved code
        		os.Remove(filePath)                                      // moved code
        ```
        * comment the :
            ```go
        		rm := download.RemoveIndexZeroDetailVideo(videoList)
        		download.WriteJSON(rm)
        		os.Remove(filePath)
            ```
    Or put it outside the loop, and make sure you get access or data from `videoList` and `filePath`,
    simply put, move the required variables outside the loop.
    * example
      * before looping
      ```go
        videoList := download.OpenDetailVideo()  // moved code
    	detailVideo := videoList[0]                // moved code
    	filePath := detailVideo.NameVideo          // moved code
    	for _, list := range listOfToken.Data {
    		// create a session from each page...
      ```
    
      * after looping
      ```go
        for _, list := range listOfToken.Data {
            // create a session from each page...
            // Upload Video...
            // published... 
    
    		publishParams.Publish()
    	}
    	rm := download.RemoveIndexZeroDetailVideo(videoList)  // moved code
    	download.WriteJSON(rm)                                // moved code
    	os.Remove(filePath)                                   // moved code
      }
      ``` 

* ### `download.DownlaodContent()` it will work to update the DATASET in APIFY and then get it as `urls.json`.
    Then continue to download all the videos and the video details will be saved as `detail_video.json`. If you don't want to download the video directly, you don't need to run this function, you just need to call the `geturl.GetDatasetFromAPIFY()` function function.
* ### `geturl.NewDatasetToAPIFY()` If you just want to update or add a DATASET, you can use this function. and you can see the results at [Storage Apify](https://console.apify.com/storage/datasets).












