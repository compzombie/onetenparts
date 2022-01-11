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

import FileChooser from './FileChooser.js'



function OptionsBar(props) {

    const [open, setOpen] = React.useState(false)

    const getFileName = () => {

    }

    const handleLoadFile = (e) => {

    };

    /*
    const handleDrawerClose = () => {
        setOpen(false)
    };
    */


    return (
        <Box sx={{ flexGrow: 1 }}>
            <AppBar position="static">
                <Toolbar>
                    <FileChooser />
                    <Button variant="contained" onClick={handleLoadFile}>Load</Button> 
                    <InputLabel>Filter:</InputLabel>
                        <Select defaultValue={0}>
                            <MenuItem value={0}>States</MenuItem>
                            <MenuItem value={1}>Binary</MenuItem>
                            <MenuItem value={2}>Domain</MenuItem>
                        </Select>
                </Toolbar>
            </AppBar>
        </Box>
    );
}

export default OptionsBar