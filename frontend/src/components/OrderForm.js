import React, { useState } from 'react';
import { handleOrder } from '../api';

const OrderForm = () => {
    const [orderSize, setOrderSize] = useState('');
    const [result, setResult] = useState(null);

    const submitOrder = async (e) => {
        e.preventDefault();
        try {
            const response = await handleOrder(Number(orderSize));
            setResult(response);
        } catch (error) {
            console.error('Failed to submit order:', error);
        }
    };

    return (
        <div>
            <h2>Order Packages</h2>
            <form onSubmit={submitOrder}>
                <input
                    type="number"
                    value={orderSize}
                    onChange={(e) => setOrderSize(e.target.value)}
                    placeholder="Enter order size"
                />
                <button type="submit">Submit</button>
            </form>
            {result && (
                <div>
                    <h3>Package Breakdown</h3>
                    <ul>
                        {Object.keys(result).map((size) => (
                            <li key={size}>
                                Package Size: {size}, Quantity: {result[size]}
                            </li>
                        ))}
                    </ul>
                </div>
            )}
        </div>
    );
};

export default OrderForm;
