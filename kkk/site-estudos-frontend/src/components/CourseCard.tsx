import React from 'react';

interface CourseCardProps {
    title: string;
    description: string;
    image: string;
}

const CourseCard: React.FC<CourseCardProps> = ({ title, description, image }) => {
    return (
        <div className="course-card">
            <img src={image} alt={title} className="course-card-image" />
            <h3 className="course-card-title">{title}</h3>
            <p className="course-card-description">{description}</p>
        </div>
    );
};

export default CourseCard;