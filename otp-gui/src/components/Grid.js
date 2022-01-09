
import { TableBody, TableContainer, TableHead } from '@mui/material';
import React, { useState } from 'react';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import Table from '@mui/material/Table';
import TableCell from '@mui/material/TableCell';

import Cell from './Cell.js'

function Grid(props) {

    return (
        <TableContainer component={Paper}>
            <Table>
                <TableBody>
                    <TableRow>
                        <TableCell>
                            <Cell initState="0"/>
                        </TableCell>
                    </TableRow>
                </TableBody>
            </Table>
        </TableContainer>
    );
}

export default Grid