 # Frontend notes

- A piece that has scores has to look different from a piece that doesn't. This is definitely a decision for the frontend. The backend only gives a piece that is acousmatic, for example, and the frontend needs to know "this means it won't have a score"

- The display part should be almost the same for admins and public, with these differences: scores shown to admins in the foreground should be full, scores shown to public are previews. Admins have more buttons, so the UI has to be clean enough for it to seem like things belong on both of them (hence my idea of a minimalistic UI - not material, that's too corporate)

- The homepage is the picture + bio. If there's nothing to show on the homepage for public, there's a default "Under construction" page. Admins get a setup screen

- There is a piece view that has a list of pieces and there is an event view that I would like to show somehow on a timeline, but I'm not too sure how. These are toggleable at all times, maybe on a navbar

- The event view should show mosaics of pieces, although this might be very empty because, although the system supports various pieces in the same event, this is so unlikely that I should make one piece at a time stand out. Same thing for piece view, the main attractions are audio and pdf, or just audio, with a small timeline/calendar (a calendar might be overkill and doesn't show a good picture) to show the events - With time, I would like to make loading appear like ghost components, but that's way overkill for now, I believe
