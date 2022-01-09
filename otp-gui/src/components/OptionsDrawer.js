import { Drawer, Toolbar } from '@mui/material';
import React, { useState } from 'react';
import MuiAppBar from '@mui/material/AppBar';
import IconButton from '@mui/material/IconButton';


function OptionsDrawer(props) {

    const [open, setOpen] = React.useState(false)
    const handleDrawerOpen = () => {
        setOpen(true)
    };
    const handleDrawerClose = () => {
        setOpen(false)
    };

    return (
        <div>
            <Toolbar>
                <IconButton onClick={handleDrawerOpen}>

                </IconButton>
            </Toolbar>

            <Drawer variant="persistent" anchor="left" open={open}>
                    <IconButton onClick={handleDrawerClose}>
                    </IconButton>
            </Drawer> 
        </div>
    );
}

export default OptionsDrawer