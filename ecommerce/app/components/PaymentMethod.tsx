interface Props {
    value: string;
    onChange(value: string): void;
}

export default function PaymentMethod({
    value,
    onChange
}: Props) {

    return (
        <div>

            <h3>Payment Method</h3>

            <label>
                <input
                    type="radio"
                    checked={value === "Credit Card"}
                    onChange={() =>
                        onChange("Credit Card")
                    }
                />
                Credit Card
            </label>

            <br />

            <label>
                <input
                    type="radio"
                    checked={value === "PayPal"}
                    onChange={() =>
                        onChange("PayPal")
                    }
                />
                PayPal
            </label>

        </div>
    );
}