import React from 'react'
import PropTypes from 'prop-types'
import TextField from 'material-ui/TextField';

class Search extends React.Component{
    constructor(){
        super();
        this.onSearch = this.onSearch.bind(this);
        this.state = {
            searchTerm: "", 
            delayTimer: {}
        };

        this.onSearch = this.onSearch.bind(this);
        this.onChange = this.onChange.bind(this);
    }

    onSearch(){
        this.props.onSearch(this.state.searchTerm);
    }

    onChange(event, newValue){
        clearTimeout(this.state.delayTimer);
        var delayTimer = setTimeout(this.onSearch, 300);
        this.setState({
            delayTimer: delayTimer,
            searchTerm: newValue
        });
    }

    render(){
        return (
            <TextField
                hintText="Search"
                onChange={this.onChange}
                />
        )
    }
}

Search.propTypes = {
    onSearch: PropTypes.func.isRequired
}


module.exports = Search