import React, { useCallback, useState } from 'react';

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

const revenueStyle = {
    paddingLeft: '20px',
}


export const Header = () => {
    const [revTotal, setRevTotal] = useState(0)
    const onDrop = useCallback(acceptedFiles => {
        const fd = new FormData()
        fd.append('file', acceptedFiles[0])
        fetch('http://localhost:8080/upload', {
            headers: new Headers({
                 'Authorization': 'Bearer shmoken', 
         }),
            body: fd,
            method: 'POST',
        }).then(res => res.json())
        .then(data => {
            fetch('http://localhost:8080/revenue', {
                headers: new Headers({
                    'Authorization': 'Bearer shmoken', 
                })
            }).then(res => res.json())
                .then(data => {
                    setRevTotal(data)
                })
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
            <div style={revenueStyle}>
                Total sales revenue: { revTotal }
            </div>
        </span>
    );
}
