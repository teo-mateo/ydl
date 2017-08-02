import React from 'react'
import Search from './Search'
import PropTypes from 'prop-types'

class Header extends React.Component{
    constructor(){
        super()
    }



    render(){

        return (
            <div className="flex-container">
                <div className="flex-item">Get your fix.</div>
                <div className="flex-item"><Search onSearch={this.props.onSearch} /></div>
            </div>
        )
    }
}

Header.propTypes = {
    onSearch: PropTypes.func.isRequired
}

module.exports = Header