var axios = require('axios');
var path = require('path')

var Api = {
    BaseURL: "localhost:8080", 
    GetUri: function(p, ...params){
        var uri = path.join.apply(null, [Api.BaseURL, p].concat(params))
        return "http://" + window.encodeURI(uri)
    }

}

Api.GetList = function(who){
    var uri = Api.GetUri('/list/json/', who)
    return axios.get(uri)
        .then(function(response){
            return response.data
        });
}

Api.GetUsers = function(){
    var uri = Api.GetUri('/users')
    return axios.get(uri)
        .then(function(response){
            return response.data
        });
}


module.exports = Api;