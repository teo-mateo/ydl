var axios = require('axios');
var path = require('path')
var config = require('../../ydl.config.js')

var Api = {
    GetUri: function(p, ...params){
        var uri = path.join.apply(null, [config.baseydlurl, p].concat(params))
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
            return response.data;
        })
        .catch(function(err){
            console.log("Axios: error in Api.GetUsers(): " + err);
        });
}

Api.MultiDelete = function(ids){
    var uri = Api.GetUri('/multidelete');
    return axios.post(uri, ids)
        .then(function(response){
            return response.data;
        })
        .catch(function(err){
            console.log("Axios: error in Api.MultiDelete(): " + err);
        })
}


module.exports = Api;