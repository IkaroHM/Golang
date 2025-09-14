import React from 'react';
import CourseCard from './CourseCard';

interface Course {
    id: number;
    title: string;
    description: string;
    image: string;
}

interface RecommendedCoursesProps {
    courses: Course[];
}

const RecommendedCourses: React.FC<RecommendedCoursesProps> = ({ courses }) => {
    return (
        <div className="recommended-courses">
            <h2>Recommended Courses</h2>
            <div className="course-list">
                {courses.map(course => (
                    <CourseCard 
                        key={course.id} 
                        title={course.title} 
                        description={course.description} 
                        image={course.image} 
                    />
                ))}
            </div>
        </div>
    );
};

export default RecommendedCourses;