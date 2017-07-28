import React from 'react'
import IconButton from 'material-ui/IconButton';
import FileCloudDownload from 'material-ui/svg-icons/file/cloud-download';
import ActionsDelete from 'material-ui/svg-icons/action/delete';

import PropTypes from 'prop-types';

class MultiActions extends React.Component{
    constructor(){
        super()
    }

    render(){
        var hasSelection = false;
        if(this.props.selectedSongs.length > 0){
            hasSelection = true;
        }

        return (
            <div>
                <span>With selected: </span>
                <a href='#'>
                    <IconButton disabled={hasSelection ? false : true}>
                        <FileCloudDownload />
                    </IconButton>
                </a>                                    
                <a href='#'>
                    <IconButton disabled={hasSelection ? false : true}>
                        <ActionsDelete />
                    </IconButton>
                </a>                                                 
            </div>    
        )
    }
}

MultiActions.propTypes = {
    selectedSongs: PropTypes.array.isRequired
}

module.exports = MultiActions;