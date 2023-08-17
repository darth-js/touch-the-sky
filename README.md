# Videos API
Simple restful API to manage videos and related annotations

## Context
### Main Requirements 
- Create a simple restful API to manage videos and related annotations
- Each video may have many annotations related to it
  - An annotation allows us to capture time-related information about the video. For example, we may want to create an annotation that references a part of the video between `00:04:00` to `00:05:00` that contains an advertisement.
- Uses a form of persistent storage such as SQLite or similar to store your data
- Add authorisation in the form of an API key or JWT token to your application
- State your assumptions where there is ambiguity in the job to be done section, and provide relevant documentation for us to evaluate your work
- Assume videos are stored as a link to a cloud storage solution or a video hosting service such as YouTube or Vimeo
- Provide a docker image to run your solution

### User Stories
- As a user, I would like to be able to create a video on the system with relevant video metadata
- As a user, I would like to be able to create an annotation with the start and end times of the annotation
- As a user, I would like the system to error if I create an annotation that is out of bounds of the duration of the video
- As a user, I would like to be able to list all annotations relevant to a video
- As a user, I would like to be able to specify annotation type and add additional notes
- As a user, I would like to be able to update my annotation details
- As a user, I would like to be able to delete my annotations
- As a user, I would like to be able to delete videos from the system

### ADRs
- [001 Selection of Architecture](docs/adr/adr-001-selection-of-architecture-design.md)
- [002 File structure](docs/adr/adr-002-file-structure.md)

## Running Application
```bash
docker compose up --build
```


