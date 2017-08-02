import React from 'react';
import PropTypes from 'prop-types';

class Pagination extends React.Component{
    constructor(){
        super()
    }

    onPageChange(page){
        this.props.onPageChange(page);
    }

    render(){
        var numberOfPages = Number.parseInt((this.props.itemsCount/this.props.pageSize).toFixed(0))
        var pages = [];
        for (var i = 0; i < numberOfPages; i++) {
            pages.push(0);
        }
        return (
            <div className='flex-container'>
                {pages.map((p, index) => {
                    return (
                        <button key={index} onClick={function(){
                            this.onPageChange.bind(this)(index+1);
                        }.bind(this)}>{index+1}</button>
                    );
                })}
            </div>
        );
    }
}

Pagination.propTypes = {
    pageSize: PropTypes.number.isRequired,
    itemsCount: PropTypes.number.isRequired,
    onPageChange: PropTypes.func.isRequired
}

module.exports = Pagination