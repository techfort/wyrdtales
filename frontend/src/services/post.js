import { get, post } from 'axios';

const BASE_URL = "http://localhost:1323/";

const URLs = {
    latest: `${BASE_URL}posts/latest`,
    saveStory: `${BASE_URL}posts/story`,
    storyById: id => `${BASE_URL}posts/story/${id}`,
    blogById: id => `${BASE_URL}posts/blogpost/${id}`
};

const getLatest = () => {
    return get(URLs.latest);
};

const getStory = id => {
    return get(URL.storyById(id));
};

const getBlogPost = id => {
    return get(URL.blogById(id));
};

const saveStory = story => {
    return post(URL.saveStory, story);
};

export default {
    getLatest,
    getStory,
    getBlogPost,
    saveStory
};