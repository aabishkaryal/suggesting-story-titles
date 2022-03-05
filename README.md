

# Backend Engineer Tech Task

## Suggesting Story Titles

**Background**

Popsa helps people rediscover and relive their best experiences and stories. Users create photo-personalised products, such as photobooks, using our apps. Our aim is for this storytelling process to be as simple as possible - we aim to simplify and automate as much as possible; preferring intelligent suggestions over a plethora of buttons and tools.

This philosophy has led us to reduce our users’ average journey time to 6 minutes (for context, incumbents in our industry benchmark their journey time in hours or days).

**The Task**

Every good story needs a title and, with simplicity in mind, we would like you to suggest a title or titles for a given group of photos.

For this task we have prepared 3 CSV ﬁles, each ﬁle contains a list of photo-metadata and the ﬁle represents one “album” of photos. The metadata consists of:
- A timestamp (representing the date the photo was taken)
- Geo-coordinates (latitude and longitude)

**Please design and implement a component that examines each ﬁle and suggests appropriate titles or a title.**

### Example Data:

| **Date Photo Was Taken** | **Latitude** | **Longitude** |
| ------------------------ | ------------ | ------------- |
| 2019-03-30 14:12:19      | 40.703717    | -74.016094    |
| 2019-03-30 15:34:49      | 40.782222    | -73.965278    |
| 2019-03-31 12:18:04      | 40.748433    | -73.985656    |

**Example Output**

- “A weekend in New York”
- “A trip to New York”
- “A weekend in Manhattan”
- “A rainy trip to New York”
- “A trip to the United States”
- “New York in March”

### Data Set

The data-set is available here:
<https://gist.github.com/tomjcohen/726d24f1fe2736a16028911c3b544bfc>

(This is provided to test/demo your solution but you are welcome to create your own...)

Task Details

1. Bearing in mind that titles will be displayed to users, credit is given for creative output, try and tailor output to give the most appropriate suggestions

1. Integrate with a third-party API for reverse-geocoding (Google Maps, Here.com both provide free tiers)s