document.addEventListener('DOMContentLoaded', function() {
    // DOM Elements
    const userForm = document.getElementById('userForm');
    const registrationForm = document.getElementById('registrationForm');
    const votingSection = document.getElementById('votingSection');
    const thankYouSection = document.getElementById('thankYouSection');
    const resultsSection = document.getElementById('resultsSection');
    const voteButtons = document.querySelectorAll('.vote-btn');
    const catCount = document.getElementById('catCount');
    const dogCount = document.getElementById('dogCount');
    const totalVotes = document.getElementById('totalVotes');

    // Chart initialization
    const ctx = document.getElementById('voteChart').getContext('2d');
    let voteChart;

    // API URLs
    const API_BASE_URL = '/api';
    const VOTE_ENDPOINT = `${API_BASE_URL}/vote`;
    const STATS_ENDPOINT = `${API_BASE_URL}/stats`;

    // User data
    let userData = {
        name: '',
        email: ''
    };

    // Initialize chart
    initChart();

    // Load initial stats
    fetchStats();

    // Set up polling for stats
    setInterval(fetchStats, 5000);

    // Event Listeners
    userForm.addEventListener('submit', function(e) {
        e.preventDefault();

        // Get user data
        userData.name = document.getElementById('name').value.trim();
        userData.email = document.getElementById('email').value.trim();

        if (validateForm(userData)) {
            // Show voting section
            registrationForm.classList.add('hidden');
            votingSection.classList.remove('hidden');

            // Animate entrance
            votingSection.style.animation = 'fadeIn 0.5s ease forwards';
        }
    });

    voteButtons.forEach(button => {
        button.addEventListener('click', function() {
            const choice = this.getAttribute('data-choice');
            submitVote(choice);
        });
    });

    // Functions
    function validateForm(data) {
        // Basic validation
        if (!data.name) {
            showError('Please enter your name');
            return false;
        }

        if (!data.email || !isValidEmail(data.email)) {
            showError('Please enter a valid email address');
            return false;
        }

        return true;
    }

    function isValidEmail(email) {
        const re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
        return re.test(String(email).toLowerCase());
    }

    function showError(message) {
        alert(message);
    }

    function submitVote(choice) {
        // Prepare data
        const voteData = {
            name: userData.name,
            email: userData.email,
            choice: choice
        };

        // Send vote to API
        fetch(VOTE_ENDPOINT, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(voteData)
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Failed to submit vote');
                }
                return response.json();
            })
            .then(data => {
                // Show thank you message
                votingSection.classList.add('hidden');
                thankYouSection.classList.remove('hidden');

                // Update stats
                fetchStats();

                // After 3 seconds, hide thank you and show results
                setTimeout(() => {
                    thankYouSection.classList.add('hidden');
                }, 3000);
            })
            .catch(error => {
                showError('Error submitting your vote. Please try again.');
                console.error('Error:', error);
            });
    }

    function fetchStats() {
        fetch(STATS_ENDPOINT)
            .then(response => response.json())
            .then(data => {
                updateStats(data);
            })
            .catch(error => {
                console.error('Error fetching stats:', error);
            });
    }

    function updateStats(data) {
        // Update counters
        catCount.textContent = data.cats || 0;
        dogCount.textContent = data.dogs || 0;
        totalVotes.textContent = (data.cats || 0) + (data.dogs || 0);

        // Update chart
        updateChart(data);
    }

    function initChart() {
        voteChart = new Chart(ctx, {
            type: 'doughnut',
            data: {
                labels: ['Cats', 'Dogs'],
                datasets: [{
                    data: [0, 0],
                    backgroundColor: [
                        getComputedStyle(document.documentElement).getPropertyValue('--cat-color'),
                        getComputedStyle(document.documentElement).getPropertyValue('--dog-color')
                    ],
                    borderWidth: 0
                }]
            },
            options: {
                responsive: true,
                cutout: '70%',
                plugins: {
                    legend: {
                        position: 'bottom',
                        labels: {
                            font: {
                                family: 'Poppins',
                                size: 14
                            },
                            padding: 20
                        }
                    },
                    tooltip: {
                        callbacks: {
                            label: function(context) {
                                const total = context.dataset.data.reduce((a, b) => a + b, 0);
                                const value = context.raw;
                                const percentage = total > 0 ? Math.round((value / total) * 100) : 0;
                                return `${context.label}: ${value} votes (${percentage}%)`;
                            }
                        }
                    }
                },
                animation: {
                    animateScale: true,
                    animateRotate: true
                }
            }
        });
    }

    function updateChart(data) {
        voteChart.data.datasets[0].data = [data.cats || 0, data.dogs || 0];
        voteChart.update();
    }
});