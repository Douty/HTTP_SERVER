document.getElementById('testBtn').addEventListener('click', () => {
    const msg = document.getElementById('msg');
    const now = new Date().toLocaleTimeString();
    
    msg.innerText = `JS execution successful at ${now}!`;
    msg.style.color = '#e67e22';
    
    console.log("Server test button clicked.");
});