import React, { useCallback } from 'react';

import { useDropzone } from 'react-dropzone';
import csv from 'csv';

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
       const reader = new FileReader();
       reader.onabort = () => console.log("file reading was aborted");
       reader.onerror = () => console.log("file reading failed");
       reader.onload = () => {
           csv.parse(reader.result, (_, data) => {
               console.log("Parsed CSV data: ", data);
           });
       };

       acceptedFiles.forEach(file => reader.readAsBinaryString(file));
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
