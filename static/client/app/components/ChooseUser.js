import React from 'react'
import Api from '../utils/Api'

import DropDownMenu from 'material-ui/DropDownMenu';
import MenuItem from 'material-ui/MenuItem';
import PropTypes from 'prop-types';

const styles = {
  customWidth: {
    width: 200,
  },
};

class ChooseUser extends React.Component{
    constructor(){
        super()
        this.state = {
            selectedUser: "",
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
        console.log("in handleMenuChange(), value is " + (value === "" ? "<All>" : value))
        this.setState({
            selectedUser: value
        });
        this.props.onSelect(value);
    }

    render(){
        return (
                <DropDownMenu value={this.state.selectedUser} onChange={this.handleMenuChange.bind(this)}>
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
    onSelect: PropTypes.func.isRequired
}

module.exports = ChooseUser