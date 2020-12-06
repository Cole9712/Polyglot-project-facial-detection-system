const axios = require('axios');
const querystring = require('querystring');
const apiURL = 'https://pixabay.com/api/';
const apiKey = '19068932-84221674dc0aa1b3c5df75bb2';

export const photosMix = {
    methods: {
        getPhotos(page = 1) {
            const param = {
                page,
                key : apiKey,
                per_page: 21,
            }
            const queryString = querystring.stringify(param);
            return axios.get(`${apiURL}/?${queryString}`);
        },
        searchPhoto(data) {
            let param = Object.assign({}, data);
            param['key'] = apiKey;
            param['per_page'] = 90;
            // clean empty keys
            Object.keys(param).forEach(key => {
                if (!param[key]) {
                    delete param[key];
                }
            })
            const queryString = querystring.stringify(param);
            return axios.get(`${apiURL}/?${queryString}`);
        }
    }
}