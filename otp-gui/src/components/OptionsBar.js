import React, { useState } from 'react';
import MuiAppBar from '@mui/material/AppBar';
import IconButton from '@mui/material/IconButton';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import MenuItem from '@material-ui/core/MenuItem';
import FormHelperText from '@material-ui/core/FormHelperText';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';

import InputLabel from '@mui/material/InputLabel';


function OptionsBar(props) {

    const [open, setOpen] = React.useState(false)
    const handleDrawerOpen = () => {
        setOpen(true)
    };
    const handleDrawerClose = () => {
        setOpen(false)
    };


    return (
        <Box sx={{ flexGrow: 1 }}>
            <AppBar position="static">
                <Toolbar>
                    <input type="file" />
                    <InputLabel>Filter:</InputLabel>
                        <Select >
                            <MenuItem value={0}>States</MenuItem>
                            <MenuItem value={1}>Binary</MenuItem>
                            <MenuItem value={2}>Domain</MenuItem>
                        </Select>
                    {/* 
                    <Select
                        labelId="demo-simple-select-label"
                        id="demo-simple-select"
                        value={open}
                        onChange={setOpen(!open)}
                    >
                        <MenuItem value={10}>Ten</MenuItem>
                        <MenuItem value={20}>Twenty</MenuItem>
                        <MenuItem value={30}>Thirty</MenuItem>
                    </Select>
                    */}
                </Toolbar>
            </AppBar>
        </Box>
    );
}

export default OptionsBar