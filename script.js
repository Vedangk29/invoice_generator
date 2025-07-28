document.addEventListener('DOMContentLoaded', () => {
    const addItemBtn = document.getElementById('addItem');
    const itemsContainer = document.getElementById('items');
    const generateInvoiceBtn = document.getElementById('generateInvoice');

    addItemBtn.addEventListener('click', () => {
        const itemDiv = document.createElement('div');
        itemDiv.className = 'item';
        itemDiv.innerHTML = `
            <div class="form-group">
                <label>Description:</label>
                <input type="text" class="item-desc" required>
            </div>
            <div class="form-group">
                <label>Quantity:</label>
                <input type="number" class="item-qty" min="1" required>
            </div>
            <div class="form-group">
                <label>Unit Price:</label>
                <input type="number" class="item-price" min="0" step="0.01" required>
            </div>
            <button class="remove-item">Remove</button>
        `;
        itemsContainer.appendChild(itemDiv);
    });

    itemsContainer.addEventListener('click', (e) => {
        if (e.target.classList.contains('remove-item')) {
            if (itemsContainer.children.length > 1) {
                e.target.parentElement.remove();
            }
        }
    });

    generateInvoiceBtn.addEventListener('click', async () => {
        const customerName = document.getElementById('customerName').value;
        const discount = parseFloat(document.getElementById('discount').value) || 0;
        const taxRate = parseFloat(document.getElementById('taxRate').value) || 0;
        const items = Array.from(document.querySelectorAll('.item')).map(item => ({
            description: item.querySelector('.item-desc').value,
            quantity: parseInt(item.querySelector('.item-qty').value),
            unitPrice: parseFloat(item.querySelector('.item-price').value)
        }));

        if (!customerName || items.some(item => !item.description || !item.quantity || !item.unitPrice)) {
            alert('Please fill in all required fields');
            return;
        }

        const invoice = {
            customerName,
            items,
            discount,
            taxRate
        };

        try {
            const response = await fetch('/generate', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(invoice)
            });

            if (!response.ok) {
                throw new Error('Failed to generate invoice');
            }

            const blob = await response.blob();
            const url = window.URL.createObjectURL(blob);
            const a = document.createElement('a');
            a.href = url;
            a.download = 'invoice.pdf';
            document.body.appendChild(a);
            a.click();
            a.remove();
            window.URL.revokeObjectURL(url);
        } catch (error) {
            alert('Error generating invoice: ' + error.message);
        }
    });
});