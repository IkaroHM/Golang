import React from 'react';
import SearchBar from '../components/SearchBar';
import RecommendedCourses from '../components/RecommendedCourses';

const Home: React.FC = () => {
    return (
        <div>
            <h1>Bem-vindo ao Site de Estudos</h1>
            <SearchBar />
            <RecommendedCourses />
        </div>
    );
};

export default Home;