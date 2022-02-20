const mongoose = require('mongoose');

const fridgeSchema = new mongoose.Schema({
    _id: Number,
    food: {
        type: Array,
    },
    createdAt: {
       type: Date,
       default: Date.now,
    } 
}, {_id: false});

module.exports = mongoose.model('Fridge', fridgeSchema);