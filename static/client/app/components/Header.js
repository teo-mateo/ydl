import React from 'react'
import Search from './Search'

class Header extends React.Component{
    constructor(){
        super()
    }

    render(){

        return (
            <div className="flex-container">
                <div className="flex-item">Get your fix.</div>
                <div className="flex-item"><Search /></div>
            </div>
        )
    }
}

module.exports = Header