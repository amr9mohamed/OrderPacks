import React, { useState } from 'react';
import { setPackageSizes } from '../api';

const PackageSizeForm = () => {
    const [sizes, setSizes] = useState('');

    const submitSizes = async (e) => {
        e.preventDefault();
        const newSizes = sizes.split(',').map(Number);
        try {
            await setPackageSizes(newSizes);
            alert('Package sizes updated successfully');
        } catch (error) {
            console.error('Failed to update package sizes:', error);
        }
    };

    return (
        <div>
            <h2>Set Package Sizes</h2>
            <form onSubmit={submitSizes}>
                <input
                    type="text"
                    value={sizes}
                    onChange={(e) => setSizes(e.target.value)}
                    placeholder="Enter sizes (comma separated)"
                />
                <button type="submit">Update Sizes</button>
            </form>
        </div>
    );
};

export default PackageSizeForm;
