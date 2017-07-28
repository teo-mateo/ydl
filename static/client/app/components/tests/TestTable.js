import React from 'react'


import {
    Table,
    TableBody,
    TableHeader,
    TableHeaderColumn,
    TableRow,
    TableRowColumn
} from 'material-ui/Table'


class TestTable extends React.Component{
    constructor(){
        super()
        this.state = {
            people: [{
                id: 1, 
                name: 'john smith',
                status: 'employed',
                selected: true
            }, {
                id: 2,
                name: 'randal white',
                status: 'employed',
                selected: true
            }, {
                id: 3,
                name: 'stephanie sanders',
                status: 'employed', 
                selected: false
            }]
        }
    }

    render(){
        return (
            <Table selectable={true} multiSelectable={true}>
            <TableHeader>
                <TableRow>
                    <TableHeaderColumn>ID</TableHeaderColumn>
                    <TableHeaderColumn>Name</TableHeaderColumn>
                    <TableHeaderColumn>Status</TableHeaderColumn>
                </TableRow>
            </TableHeader>
            <TableBody>
                {
                    this.state.people.map(p =>{
                        return (
                            <TableRow selectable={true} key={p.id}>
                                <TableRowColumn>{p.id}</TableRowColumn>
                                <TableRowColumn>{p.name}</TableRowColumn>
                                <TableRowColumn>{p.status}</TableRowColumn>
                            </TableRow>
                        );
                    })
                }
            </TableBody>
        </Table>
        )
    }
}

module.exports = TestTable;