import React from 'react';
import { Link } from 'react-router-dom';

const FillGaps = () => {
    return (
        <div className="under-construction">
            <h1>This page is under construction</h1>
            <p>
                <Link to="/mainpage">Return to homepage</Link>
            </p>
        </div>
    );
};

export default FillGaps;
