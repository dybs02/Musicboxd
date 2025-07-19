

const extractSearchResults = (type: string, result: any) => {
  let search_res = []

  if (type === 'Track') {
    search_res = (result?.search?.tracks?.items ?? []).map((track: any) => {
      return {
        name: track.name,
        id: track.id,
        artists: track.artists,
        images: track.album.images
      }
    })
  }
  
  if (type === 'Album') {
    console.log(result)
    search_res = result?.search?.albums?.items ?? []
  }

  return search_res
}


export {
  extractSearchResults,
}