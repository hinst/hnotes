import React from 'react';

class NoteList extends React.Component {

	render() {
		const notes = this.props.notes.map(
			(note) => {
				return (<li key={note.Title}>{note.Title}</li>);
			}
		);
		return (
			<ul className="w3-ul w3-hoverable">
				{notes}
			</ul>
		);
	}

}

export default NoteList;