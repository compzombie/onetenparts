import React, { useState } from 'react';
import TextField from '@mui/material/TextField';

function Cell(props) {
    const [cellState, setCellState] = useState(props.initState)

    return (
        <div>
            <TextField label={cellState} variant="outlined"/>
        </div>
    );
}

export default Cell