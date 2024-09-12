import axios from 'axios';

// Set the base URL for your API
const API_BASE_URL = 'http://localhost:8080'; // Update with your actual API URL

// Function to handle order size and get the package breakdown
export const handleOrder = async (orderSize) => {
    try {
        const response = await axios.get(`${API_BASE_URL}/handleOrder/${orderSize}`);
        return response.data;
    } catch (error) {
        console.error('Error handling order:', error);
        throw error;
    }
};

// Function to set new package sizes
export const setPackageSizes = async (newSizes) => {
    try {
        const response = await axios.post(`${API_BASE_URL}/setPackageSizes`, { "sizes":newSizes });
        return response.data;
    } catch (error) {
        console.error('Error setting package sizes:', error);
        throw error;
    }
};
