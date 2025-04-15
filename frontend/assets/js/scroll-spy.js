/**
 * Scroll Spy for Table of Contents
 * 
 * This script adds active state to TOC links based on scroll position
 */
document.addEventListener('DOMContentLoaded', function() {
  // Initialize scrollspy both on page load and when content changes (HTMX)
  // Small delay to ensure HTMX has time to start loading the TOC
  setTimeout(initScrollSpy, 100);
  
  // Also listen for HTMX content swap events to reinitialize when navigating between blogs
  document.body.addEventListener('htmx:afterSwap', function(event) {
    // Only reinitialize if the swap target is the main content
    if (event.detail.target.id === 'content-section') {
      console.log('HTMX content swap detected, reinitializing scroll spy');
      // Give some time for content to render fully
      setTimeout(initScrollSpy, 200);
    }
  });
  
  // Specifically listen for the TOC being loaded
  document.body.addEventListener('htmx:afterSettle', function(event) {
    // If the swap target is the sidebar container where TOC is loaded
    if (event.detail.target.id === 'variable-sidebar-container') {
      console.log('TOC loaded via HTMX, initializing scroll spy');
      setTimeout(initScrollSpy, 50);
    }
  });
});

// Moved outside to be accessible from both event handlers
function initScrollSpy() {
  console.log('Initializing Scroll Spy');
  
  // Only run if we have both a table of contents and article content
  const tocLinks = document.querySelectorAll('.blog-toc .toc-list a');
  const article = document.querySelector('article');
  
  console.log('TOC Links found:', tocLinks.length);
  console.log('Article found:', !!article);
  
  if (!tocLinks.length || !article) {
    console.log('Scroll Spy waiting: Missing TOC or article');
    // No need for recursive retry - HTMX event will trigger when content is loaded
    return;
  }
    
  // Create a map of headings used in the TOC
  const tocHeadings = {};
  tocLinks.forEach(link => {
    const href = link.getAttribute('href');
    if (href && href.startsWith('#')) {
      tocHeadings[href.substring(1)] = true;
    }
  });
  
  // Only get headings that are actually in the TOC
  const headings = Array.from(article.querySelectorAll('[id]')).filter(el => 
    tocHeadings[el.id]
  );
  
  console.log('TOC Headings found:', headings.length);
  
  if (!headings.length) {
    console.log('Scroll Spy exiting: No headings with IDs found in TOC');
    return;
  }

  // Remove any previous event listeners to prevent duplicates when switching blogs
  if (window.scrollSpyActive) {
    console.log('Cleaning up previous scroll spy');
    window.removeEventListener('scroll', window.scrollSpyScrollHandler);
    window.removeEventListener('resize', window.scrollSpyResizeHandler);
  }

  // Track last active heading to prevent frequent changes
  let lastActiveHeading = null;
  // Minimum pixel distance required to change active heading
  const minScrollDistance = 30;
  // Offset from the top of the viewport for heading detection
  const topOffset = 120;
  // Create an array to store heading positions for better reference
  const headingPositions = [];

  // Pre-calculate initial heading positions and store them
  function calculateHeadingPositions() {
    headingPositions.length = 0; // Clear array
    headings.forEach(heading => {
      headingPositions.push({
        id: heading.id,
        top: heading.getBoundingClientRect().top + window.scrollY
      });
    });
    // Sort by position top to bottom
    headingPositions.sort((a, b) => a.top - b.top);
  }
  
  // Calculate initial positions
  calculateHeadingPositions();
  
  // Set a function to check which heading is in view
  function setActiveHeading() {
    const scrollPosition = window.scrollY;
    let currentHeadingId = null;
    
    // Find the current heading based on scroll position
    for (let i = headingPositions.length - 1; i >= 0; i--) {
      // If we've scrolled past this heading (with offset)
      if (scrollPosition + topOffset >= headingPositions[i].top) {
        currentHeadingId = headingPositions[i].id;
        break;
      }
    }
    
    // Special case for the first heading when near the top of the page
    if (!currentHeadingId && headingPositions.length > 0 && 
        Math.abs(scrollPosition - headingPositions[0].top) < 200) {
      currentHeadingId = headingPositions[0].id;
    }
    
    // If we found a heading ID, update the active state
    if (currentHeadingId) {
      // Find the actual heading element
      const currentHeading = document.getElementById(currentHeadingId);
      
      // Only change if it's different from the last active heading
      if (!lastActiveHeading || currentHeadingId !== lastActiveHeading.id) {
        // Remove active class from all TOC links
        tocLinks.forEach(link => {
          link.classList.remove('active');
        });
        
        // Add active class to the current heading's TOC link
        const activeLink = document.querySelector(`.blog-toc .toc-list a[href="#${currentHeadingId}"]`);
        
        if (activeLink) {
          activeLink.classList.add('active');
          // Remember this heading as the last active one
          lastActiveHeading = currentHeading;
        }
      }
    }
  }
  
  // Store the scroll handler so we can remove it later if needed
  window.scrollSpyScrollHandler = function() {
    if (!window.scrollSpyAnimating) {
      window.scrollSpyAnimating = true;
      window.requestAnimationFrame(function() {
        setActiveHeading();
        window.scrollSpyAnimating = false;
      });
    }
  };
  
  // Store the resize handler so we can remove it later if needed
  window.scrollSpyResizeHandler = function() {
    calculateHeadingPositions();
    setActiveHeading();
  };
  
  // Add throttled scroll event listener for better performance
  window.addEventListener('scroll', window.scrollSpyScrollHandler);
  
  // Add resize listener
  window.addEventListener('resize', window.scrollSpyResizeHandler);
  
  // Mark that scroll spy is active
  window.scrollSpyActive = true;
  
  // Update heading positions when images or other resources finish loading
  window.addEventListener('load', function() {
    calculateHeadingPositions();
    setActiveHeading();
  });
  
  // Run once on load with a slight delay to ensure content is fully rendered
  setTimeout(function() {
    calculateHeadingPositions();
    setActiveHeading();
    console.log('Scroll Spy initialized successfully');
  }, 300);
}
