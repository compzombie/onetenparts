import React, { useState } from 'react';

function FileChooser(props) {

    const [file, setFile] = React.useState("")

    const fileIn = React.createRef()



    return (
        <>
            <input id="upload" type="file" ref={fileIn}/>
            {fileIn.value}
        </>
    );
}

export default FileChooser