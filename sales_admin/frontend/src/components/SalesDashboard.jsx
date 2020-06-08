import React from 'react';

import { Header } from './Header';

const headerStyle = {
    padding: '10px',
    height: '30px',
    backgroundColor: 'lightblue',
}

export const SalesDashboard = () => {

    return (
        <div style={headerStyle}>
            <Header />
        </div>    
    );
}
