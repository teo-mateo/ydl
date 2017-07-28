var axios = require('axios');

module.exports = {
    fetchPopularRepos: function(language){
        var uri = 
            'https://api.github.com/search/repositories?q=stars:>1+language:' + language + 
            '&sort=stars&order=desc&type=Repositories';
        var encodeURI = window.encodeURI(uri);

        return axios.get(encodeURI)
            .then(function(response){
                return response.data.items
            })
    }
}