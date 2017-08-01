import React from 'react'
import Api from '../utils/Api'

import DropDownMenu from 'material-ui/DropDownMenu';
import MenuItem from 'material-ui/MenuItem';
import PropTypes from 'prop-types';
import { withRouter } from 'react-router-dom';

const styles = {
  customWidth: {
    width: 200,
  },
};

class ChooseUser extends React.Component{
    constructor(){
        super()
        this.state = {
            users: [""]
        }
    }

    componentDidMount(){
        Api.GetUsers()
            .then(users => {
                console.log(users);
                this.setState({users: users})
            })
    }

    handleMenuChange (event, index, value) {
        this.props.history.push('/users/'+value)
    }

    render(){
        return (
                <DropDownMenu value={this.props.selectedUser} onChange={this.handleMenuChange.bind(this)}>
                    <MenuItem value="" primaryText="All users" />
                    {
                        this.state.users.map(function(u){
                            return <MenuItem value={u} primaryText={u} key={u} />;
                        })
                    }
                </DropDownMenu>
        )
    }
}

ChooseUser.propTypes = {
    selectedUser: PropTypes.string.isRequired
}

module.exports = withRouter(ChooseUser)