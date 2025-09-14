import React, { Component } from 'react';

interface SearchBarState {
    searchTerm: string;
}

class SearchBar extends Component<{}, SearchBarState> {
    constructor(props: {}) {
        super(props);
        this.state = {
            searchTerm: ''
        };
    }

    handleSearch = (event: React.ChangeEvent<HTMLInputElement>) => {
        this.setState({ searchTerm: event.target.value });
    };

    render() {
        return (
            <div className="search-bar">
                <input
                    type="text"
                    placeholder="Buscar cursos..."
                    value={this.state.searchTerm}
                    onChange={this.handleSearch}
                />
            </div>
        );
    }
}

export default SearchBar;