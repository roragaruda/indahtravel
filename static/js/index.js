function formatToRupiah(amount) {
    if (amount >= 1000000) {
        // Convert to millions and format
        return `Rp. ${(amount / 1000000).toFixed(1).replace('.', ',')}jt`;
    } else if (amount >= 1000) {
        // Convert to thousands and format
        return `Rp. ${(amount / 1000).toFixed(1).replace('.', ',')}rb`;
    }
    // Return the amount for numbers less than 1000
    return `Rp. ${amount}`;
}

function calculateDaysBetween(startDateStr, endDateStr) {
    // Parse the date strings into Date objects
    const startDate = new Date(startDateStr);
    const endDate = new Date(endDateStr);

    // Calculate the difference in time (in milliseconds)
    const differenceInTime = endDate - startDate;

    // Convert milliseconds to days
    const differenceInDays = differenceInTime / (1000 * 60 * 60 * 24);

    return differenceInDays;
}

function loadingAnimation() {
    let frames = 0;
    const framesArray = ['-', '\\', '|', '/'];

    const animationInterval = setInterval(() => {
        // Clear the previous line
        console.clear();

        // Create the loading animation
        console.log(`Loading ${framesArray[frames]}`);
        frames = (frames + 1) % framesArray.length;
    }, 100);

    // Stop the animation after a specific duration (e.g., 5000ms)
    setTimeout(() => {
        clearInterval(animationInterval);
        console.log("Loading complete!");
    }, 5000);
}

loadingAnimation();

