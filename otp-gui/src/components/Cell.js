import React, { useState } from 'react';
import TextField from '@mui/material/TextField';

function Cell(props) {
    const [cellState, setCellState] = useState(props.initState)

    return (
        <>
            <TextField label={cellState} variant="outlined"/>
        </>
    );
}

export default Cell