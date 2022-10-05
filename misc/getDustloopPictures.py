import requests
import time
import os
import glob
from pathlib import Path

url = "https://www.dustloop.com/wiki/index.php?title=Special:CargoExport&tables=MoveData_GGST&&fields=chara%2C+images%2C+hitboxes&&order+by=%60chara%60&limit=1000&format=json"
imageMetadata = "https://dustloop.com/wiki/api.php?action=query&format=json&prop=imageinfo&titles=File:{0}&iiprop=url"

def downloadPictures():
    response = requests.get(url)
    data = response.json()
    for elem in data:
        if elem['chara'] != "Anji Mito":
            break
        downloadImage(elem, 'images')
        downloadImage(elem, 'hitboxes')


def downloadImage(elem, type):
    for imageName in elem[type]:
        if (imageName == '') or Path("../website/images/" +  imageName).is_file(): #Is the image name empty? Or does it already exist?
            print("Skipped: ", imageName)
            continue
        response2 = requests.get(imageMetadata.format(imageName))
        responseImageMetadata = response2.json()
        imagePageId = responseImageMetadata['query']['pages'].keys() #We don't know the id of the page ahead of time.
        imageUrl = responseImageMetadata['query']['pages'][list(imagePageId)[0]]['imageinfo'][0]['url']
        move_image = requests.get(imageUrl).content
        with open(imageName, 'wb') as handler:
            handler.write(move_image)
        print("Downloaded: ", imageName)
        time.sleep(2) #Think about the dustloop servers.


downloadPictures()


    