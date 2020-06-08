import React, { useCallback } from 'react';

import { useDropzone } from 'react-dropzone';

const headerRowStyle = {
    display: 'flex',
    flexDirection: 'row',
    color: 'yellow',
    fontWeight: '500',
    fontSize: '16px',
}

const headerTitleStyle = {
    marginRight: '25px',
}

const inputBoxStyle = {
    color: 'darkslategray',
    backgroundColor: 'white',
    paddingRight: '3px',
    paddingLeft: '3px',
}


export const Header = () => {
    const onDrop = useCallback(acceptedFiles => {
        const fd = new FormData()
        fd.append('form', acceptedFiles[0])
        fetch('http://localhost:8080/upload/', {
            body: fd,
            method: 'POST',
        }).then(data => {
            console.log(data)
        })
    }, []); 

    const { getRootProps, getInputProps } = useDropzone({ onDrop })
   
    return (
        <span style={headerRowStyle}>
            <div style={headerTitleStyle}>ACME Cult Hero Supplies Sales Admin</div>
            <div style={inputBoxStyle} { ...getRootProps() }>
                <input { ...getInputProps() } />
                <span>Drag CSV here or click to upload</span>
            </div>
        </span>
    );
}
